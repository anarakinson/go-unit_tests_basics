package utils


func Max(arr []int) (int){
    var max int

    for _, val := range(arr) {
        if val > max {
            max = val
        }
    }

    return max
}


func MaxIdx(arr []int) (int){
    var max int
    var idx int

    for i, val := range(arr) {
        if val > max {
            max = val
            idx = i
        }
    }

    return idx
}
