; Function (parse text)
(set parse (def t (do
    (set arr (list.new))
    (set acc (dict.new (with "value" (list.new))))

    (iterate (str-iterate (concat t " ")) c (do
        (if (or (eq c " ") (eq c "\t")) (do
            (list.add arr (str-join acc.value ""))
            (dict.set acc "value" (list.new))
        ) (do
            (list.add acc.value c)
        ))
    ))

    (return arr)
)))

; Call the parser (parse text)
(print (parse "Hello world and all inside"))

; Iterate over parser's result. Will iterate every word
(iterate (parse "This is the parsed array") word (do
    (print word)
))