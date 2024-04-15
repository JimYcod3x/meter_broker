package meters

type MeterType int
	


// Electricity Meter
// 水表 (Shuǐbiǎo) - Water Meter
// IoT (Internet of Things) - IoT (Internet of Things)
// 气表 (Qìbiǎo) - Gas Meter
// 热表 (Rèbiǎo) - Heat Meter
// PV - Photovoltaic (Solar Panels)


const  (
	ElectricityMeter  MeterType =  iota + 1// 001
	WaterMeter // 010
	IoT //011
	GasMeter// 100
	HeatMeter // 101
	PV // 110
)





