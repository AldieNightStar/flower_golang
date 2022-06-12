(repeat 1000 j (do
    (repeat 1000 i (do
        (if (and
                (eq (mod i 100) 0)
                (eq (mod j 100) 0)
            ) (do
                (print i j)
        ))
    ))
))