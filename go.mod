module internal

go 1.17

require(
    internal/utils v1.0.0
)

replace (
    internal/utils => ./internal/utils
)
