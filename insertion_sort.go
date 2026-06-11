        var i, j int
        var temp JadwalSewa

        n := len(jadwalKosong)
        i = 1

        for i <= n-1 {
            j = i
            temp = jadwalKosong[j]

            for j > 0 && temp.HargaTotal < jadwalKosong[j-1].HargaTotal {
                jadwalKosong[j] = jadwalKosong[j-1]
                j--
            }

            jadwalKosong[j] = temp
            i++
        }
        fmt.Println("\nHasil Urut Harga Sewa (Insertion Sort):")
