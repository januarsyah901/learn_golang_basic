package main;

import "fmt";

func main() {

	var (
		int32 int32 = 32768;
		int64 int64 = int64(int32);
		int16 int16 = int16(int64);
	)
	fmt.Println(int16); // melebihi kapasitas  langusng balik ke yg paling bawah


}