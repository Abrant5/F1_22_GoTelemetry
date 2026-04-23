package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", ":20777")
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	fmt.Println("🏎️  Telemetria F1 22 - Estrutura Oficial")

	buf := make([]byte, 2048)
	for {
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil || n < 24 {
			continue
		}

		if buf[5] == 6 {
			
			speed := binary.LittleEndian.Uint16(buf[24:26])

		
			
			gear := int8(buf[39])

			rpm := binary.LittleEndian.Uint16(buf[40:42])

			marchaTexto := fmt.Sprintf("%d", gear)
			if gear == 0 {
				marchaTexto = "N"
			} else if gear == -1 {
				marchaTexto = "R"
			}

			bitsSteer := binary.LittleEndian.Uint32(buf[30:34])
			steer := math.Float32frombits(bitsSteer)

			fmt.Printf("\r\033[K✅ OK | Vel: %3d km/h | RPM: %5d | Marcha: %2s | Volante: %0.2f", speed, rpm, marchaTexto, steer)
		}
	}
}