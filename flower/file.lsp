(set api
    (dict
        (with "print" print)
        (with "printCool" (def t (do
            (print "-------------------------------")
            (print t)
            (print "-------------------------------")
        )))
    )
)

(api.print "Hello!")
(api.printCool "Hello sir!")