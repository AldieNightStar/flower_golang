(iterate (range 0 1000) i (do
    ; Break if number is more than 100
    (if (greater i 100) (do (break)))
    (print i)
))