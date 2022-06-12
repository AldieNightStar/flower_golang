(repeat 1000 j (do
    (repeat 1000 i (do
        (if (eq (mod i 1000) 0) (do
            (print i j)
        ))
    ))
))