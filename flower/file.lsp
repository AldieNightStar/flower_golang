(set profile
    (dict
        (with "name" "Ihor")
        (with "age" 18)
        (with "inventory" (dict
                (with "knife" 1)
                (with "water" 3)
                (with "food" 7)
            )
        )
    )
)

(print profile.inventory.food)
(dict-set profile.inventory "food" (add profile.inventory.food 1))
(print profile.inventory.food)