package hdc1080
import(
  "log"
  "time"
  "errors"
	"golang.org/x/exp/io/i2c"
)

func ReadTempHumi(d string,a int)(float32,float32,error){
// 0x40 is TI HDC1080
  b2:=[]byte{0x0,0x0}
  b4:=[]byte{0x0,0x0,0x0,0x0}
  t:=float32(0.0)
  h:=float32(0.0)
  Dev,err:=i2c.Open(&i2c.Devfs{Dev:d},a)
  defer Dev.Close()

  if err!=nil{
    log.Println("I2C BUS Error")
    return 0.0,0.0,err
  }else{

    err:=Dev.ReadReg(0xFE,b2)
    if err!=nil{
      log.Println("0:",err)
      return 0.0,0.0,err
    }else{
      vendor:=string(b2)
      //log.Println("vendor=",vendor)
      if vendor!="TI"{
        return t,h,errors.New("Not TI HDC1080")
      }
    }
  /*
  err=d.ReadReg(0x02,b2)
  if err!=nil{
    log.Println("1:",err)
  }else{
    u64:=uint64(b2[0]<<8)+uint64(b2[1])
    log.Println("Conf=",strconv.FormatUint(u64,16))
  }
  */
  // get H & T
    Dev.Write([]byte{0x00})
    for n:=0;n<10;n++{
      err=Dev.Read(b4)
      if err==nil{
        //log.Println("Data=",b4)
        t16:=int(b4[0])<<8 + int(b4[1])
        h16:=int(b4[2])<<8 + int(b4[3])
        t=165.0*(float32(t16)/65536.0)-40.0
        h=(float32(h16)/65536)*100.0
        //log.Println("Temp=",t," Humidity=",h,"%")
        break
      }
      time.Sleep(100 * time.Millisecond)
    }
    // if BREAK, err should be nil. If NOT, No Data
    return t,h,err
  }
}
