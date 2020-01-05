package hdc1080
import(
  "log"
  "time"
	"golang.org/x/exp/io/i2c"
)

func ReadTempHumi(d *i2c.Device)(float32,float32,error){
// 0x40 is TI HDC1080
  t:=float32(0.0)
  h:=float32(0.0)
  err=d.ReadReg(0xFE,b2)
  if err!=nil{
    log.Println("0:",err)
    return 0.0,0,err
  }else{
    log.Println("vendor=",string(b2))
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
  i2c1.Write([]byte{0x00})
  // time.Sleep(10 * time.Millisecond)
  b4:=[]byte{0x00,0x00,0x00,0x00}
  for n:=0;n<10;n++{
    err=i2c1.Read(b4)
    if err==nil{
      log.Println("Data=",b4)
      t16:=int(b4[0])<<8 + int(b4[1])
      h16:=int(b4[2])<<8 + int(b4[3])
      t=165.0*(float32(t16)/65536.0)-40.0
      h=(float32(h16)/65536)*100.0
      log.Println("Temp=",t," Humidity=",h,"%")
      break
    }
    time.Sleep(100 * time.Millisecond)
  }
  // if BREAK, err should be nil. If NOT, No Data 
  return t,h,err
}
