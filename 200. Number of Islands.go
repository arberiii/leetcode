func numIslands(grid [][]byte) int {
    if grid == nil || len(grid) < 1 || len(grid[0]) < 1 {
        return 0
    }
    
    rows := len(grid)
    cols := len(grid[0])

    islands := 0
    visited := make([][]bool, rows)
    
    for row := 0; row < rows; row++ {
        visited[row] = make([]bool, cols)
    }
    
    for row := 0; row < rows; row++ {
        for col:= 0; col < cols; col++ {
            if visited[row][col] {
                continue
            }
            if grid[row][col] == '0' {
                continue
            }
            
            islands++
            bfs(grid, visited, row, col)
        }
    }
    
    return islands
}

type position struct{
    row int
    col int
}

var directions = []position {
    {row: -1, col: 0},
    {row: 0, col: -1},
    {row: 0, col: 1},
    {row: 1, col: 0},
}

func bfs(grid [][]byte, visited [][]bool, row int, col int) {
    visited[row][col] = true
    queue := []position{
        {row: row, col: col},
    }
    
    for len(queue) > 0 {
        elem := queue[0]
        queue = queue[1:]
        
        for _, direction := range directions {
            nextRow := elem.row + direction.row
            nextCol := elem.col + direction.col
            
            if !isValid(grid, nextRow, nextCol) {
                continue
            }
            if visited[nextRow][nextCol] {
                continue
            }
            if grid[nextRow][nextCol] == '0' {
                continue
            }
            
            visited[nextRow][nextCol] = true
            queue = append(queue, position{row: nextRow, col: nextCol})
        }
    }
}

func isValid(grid [][]byte, row int, col int) bool {
    if row < 0 || row >= len(grid) {
        return false
    }
    if col < 0 || col >= len(grid[0]) {
        return false
    }
    
    return true
}
