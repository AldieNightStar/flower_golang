(set rec (def n allow (do
    (print (str.concat n allow))
    (if (not allow) (do (return nil)))
    (print (str.concat "REC " n))
    (rec (div n 2) false)
)))

(rec 10 true)