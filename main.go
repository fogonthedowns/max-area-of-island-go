package main

// keep track of max
// keep track of current
// loop through grid
// - reset curIsland each loop
// - implement bfs, passing pointers

func maxAreaOfIsland(grid [][]int) int {
	maxIsland := 0
	curIsland := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			curIsland = 0
			bfs(i, j, &grid, &curIsland, &maxIsland)
		}
	}
	return maxIsland
}

// bfs searches for an island
// - if it finds one, it marks it, using the dereferenced pointer
//   - sums current island
//   - if current island is greater than or equal to max
//       - modifies max
// continue to search for other parts of the island. Do a bfs of the following:
//   - above
//   - below
//   - right
//   - left
//     - each element has a bound protector

func bfs(i int, j int, world *[][]int, curIsland *int, maxIsland *int) {
	if (*world)[i][j] == 1 {
		(*world)[i][j] = 0
		*curIsland += 1
		// fmt.Printf("* cur: %v max: %v debug: %t\n", *curIsland, *maxIsland, d)
		if *curIsland >= *maxIsland {
			*maxIsland = *curIsland
		}

		// next row
		if i+1 <= len(*world)-1 {
			bfs(i+1, j, world, curIsland, maxIsland)
		}

		// prev row
		if i > 0 {
			bfs(i-1, j, world, curIsland, maxIsland)
		}

		// next col
		if j+1 <= len((*world)[i])-1 {
			bfs(i, j+1, world, curIsland, maxIsland)
		}

		// prev col
		if j > 0 {
			bfs(i, j-1, world, curIsland, maxIsland)
		}
	}
}
