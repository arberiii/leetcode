package numberOfSteps

func numberOfSteps (num int) int {
    steps := 0
    
    for {
        if num == 0 {
            return steps
        }
        if num % 2 == 0 {
            num = num / 2
            steps = steps + 1
        } else {
            num = num - 1  
            steps = steps + 1
        }
    }
    
    return steps
}
