(set profile
    (dict
        (with "name" "Ihor")
        (with "age" 18)
        (with "city" "London")
    )
)

(print (dict-get profile "name"))

(dict-set profile "career" "programmer")

(print (dict-get profile "career"))