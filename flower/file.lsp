(set block (do
    (print name)
    (return 1111)
    (print age)
    (print size)
))

(print (call block
    (with name "Ihor")
    (with age 18)
    (with size 34)
))