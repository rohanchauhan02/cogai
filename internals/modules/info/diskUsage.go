package info

import (
	"fmt"
	"syscall"

	"github.com/spf13/cobra"
)

var diskUsageCmd = &cobra.Command{
	Use:   "du",
	Short: "Get disk usage information",
	Long:  `Get disk usage information`,
	Run: func(cmd *cobra.Command, args []string) {
		total, free, err := getDiskUsage("/")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Total disk space: %.2f GB\n", float64(total)/1e9)
		fmt.Printf("Free disk space: %.2f GB\n", float64(free)/1e9)
	},
}

func getDiskUsage(path string) (total uint64, free uint64, err error) {
	var stat syscall.Statfs_t

	err = syscall.Statfs(path, &stat)
	if err != nil {
		return
	}
	// Total = block size * total number of blocks
	total = stat.Blocks * uint64(stat.Bsize)

	// Free = block size * number of free blocks
	free = stat.Bfree * uint64(stat.Bsize)

	return
}
