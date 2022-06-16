(set cnt (box 0))
(set mut (mutex))
(iterate (range 1 10) i (do
    (thread (do
        (lock mut (do
            (set cnt.value (add cnt.value 1))
            (print cnt.value)
        ))
    ))
))

(time.sleep 2000)