(set Profile
    (dict.new
        (with "name" "default")
        (with "age" 18)
        (with "getName" (def self (do
            (return self.name)
        )))
        (with "getAge" (def self (do
            (return self.age)
        )))
    )
)
(set User
    (dict.new
        (extends Profile)
        (with "name" "Ihor")
        (with "age" 18)
    )
)

(print (User.getName User))
(print (User.getAge User))
(print (dict.get User "getAge"))