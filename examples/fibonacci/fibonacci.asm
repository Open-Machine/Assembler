# Prints all fibonacci numbers until ${fibonacciNumberIndex}th

@VAR
    fibonacciNumberIndex = 8

    oneVarConst = 1
    prev = 1
    current = 1
    auxNext = 0

@CODE
    copy fibonacciNumberIndex
    sub oneVarConst
    sub oneVarConst
    store fibonacciNumberIndex

    output prev
    output current

    for:
        copy prev
        add current
        store auxNext

        copy current
        store prev

        copy auxNext
        store current
        output current

        copy fibonacciNumberIndex
        sub oneVarConst
        je end

        store fibonacciNumberIndex
        jmp for
    end:
    kill
