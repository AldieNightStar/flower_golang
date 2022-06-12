(set ff (def a (do
    (return (add a 1))
)))

(print (ff 4))