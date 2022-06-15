(set api (dict.new))

(set api.multiplied (generator pos (do
    (if (greater pos 10) (do
        (return nil)
    ))
    (return (mul pos 2))
)))

(return api)