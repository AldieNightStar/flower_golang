(set test
    (list (range 100 200))
)

(set filter
    (def arr f (do
        (set l (list))
        (iterate arr elem (do
            (if (f elem) (do
                (list-add l elem)
            ))
        ))
        (return l)
    ))
)

(set only20 (def el (do
    (return (eq (mod el 20) 0))
)))

(print (filter test only20))
