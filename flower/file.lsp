(iterate (infinite) i (do
    (if (less i 50) (do
        (continue)
    ))
    (print i)
    (if (greater i 100) (do
        (break)
    ))
))