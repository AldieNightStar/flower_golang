(set tag (def name args (do
    (set arr (list.new))
    (list.add arr "(")
    (list.add arr concat(name " "))
    (iterate args arg (do
        (list.add arr arg)
        (list.add arr " ")
    ))
    (list.add arr ")")
    (return (str.join arr ""))
)))

(return tag)