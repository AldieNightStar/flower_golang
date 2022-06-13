; Function (parse text)
(set parse (def t (do
    (set arr (list.new))
    (set acc (dict.new (with "value" (list.new))))

    (iterate (str.iterate (str.concat t " ")) c (do
        (if (or (eq c " ") (eq c "\t")) (do
            (list.add arr (str.join acc.value ""))
            (set acc.value (list.new))
        ) (do
            (list.add acc.value c)
        ))
    ))

    (return arr)
)))

; Will return: LIST [[0] = hello, [1] = world, [2] = and, [3] = all]
(print (parse "hello world and all"))