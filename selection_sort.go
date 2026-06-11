        var t JadwalSewa
        var i, j, idx_min int

        n := len(jadwalKosong)
        i = 1

        for i <= n-1 {
            idx_min = i - 1
            j = i

            for j < n {
                if jadwalKosong[idx_min].JamMulai > jadwalKosong[j].JamMulai {
                    idx_min = j
                }
                j++
            }

            t = jadwalKosong[idx_min]
            jadwalKosong[idx_min] = jadwalKosong[i-1]
            jadwalKosong[i-1] = t

            i++
        }
        fmt.Println("\nHasil Urut Jam Mulai (Selection Sort):")
