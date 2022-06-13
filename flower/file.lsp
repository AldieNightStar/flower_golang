(iterate (infinite) i (do
    (print i)
    (if (greater i 100) (do
        (break)
    ))
))