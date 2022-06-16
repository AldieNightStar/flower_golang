(set fibonacci (def n (do
    (if (less-eq n 2) (do
        (return n)
    ) (do
        (return (add
            (fibonacci (sub n 1))
            (fibonacci (sub n 2))
        ))
    ))
)))

(print (fibonacci 5))