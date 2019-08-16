package main

import (
	"flag"    
    "os"
    "github.com/mitchellh/go-fs"    
    "github.com/mitchellh/go-fs/fat"
)

// https://www.gnu.org/software/mtools/manual/mtools.html#mformat
// --------------------------------------------------------------
// mformat -f 1440                    [-f size]
//         -l HARIBOTEOS              [-v volume_label] 
//         -N 0xffffffff              [-N serial_number] 
//         -C                         [   creates the disk image file to install the MS-DOS file system on it.] 
//         -B ${03_day_harib00g_IPLB} [-B boot_sector] 
//         -i ${03_day_harib00g_OS}   [-i output image name] 
//
func main() {

    var (
        fatSize = flag.Int64("f", 1440, "FATのサイズを指定する")
        volumeLabel = flag.String("l", "HARIBOTEOS", "FATのボリュームラベルを指定する")
        serialNumber = flag.String("N", "0xffffffff", "FATのシリアルナンバーを指定する")
        // bootSector = flag.String("B", "", "FATのブートセクタを指定する")
        outputImage = flag.String("i", "", "出力先の指定をする")
    )

	// 読み書き可能, 新規作成, ファイル内容あっても切り詰め
	floppyF, _ := os.OpenFile(*outputImage, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
    floppyF.Truncate(*fatSize * 1024)
    defer floppyF.Close()
    // ブロックデバイスを構築する
    device, _ := fs.NewFileDisk(floppyF)
	// ブロックデバイスをフォーマットする
	formatConfig := &fat.SuperFloppyConfig{
		FATType: fat.FAT12,
		Label:   *volumeLabel,
		OEMName: *serialNumber,
	}

    fat.FormatSuperFloppy(device, formatConfig)
}
