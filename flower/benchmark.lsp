(set acc (box 0))

(set bench (def (do
    (iterate (range 0 1000) i (do
        (iterate (range 0 1000) j (do
            (set acc.value (add acc.value 1))
        ))
        (if (eq (mod i 100) 0) (do
            (print i)
        ))
    ))
)))

(return bench)