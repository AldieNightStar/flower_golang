(set xy (def w h (do
    (set local (dict.new (with "x" 0) (with "y" 0)))
    (return (generator i (do
        (set res (dict.new (with "x" local.x) (with "y" local.y)))
        (set local.x (add local.x 1))
        (if (greater-eq local.x w) (do
            (set local.y (add local.y 1))
            (set local.x 0)
        ))
        (if (greater-eq local.y h) (do
            (return nil)
        ))
        (return res)
    )))
)))

(iterate (xy 10 10) pos (do
    (print (str.concat pos.x "\t" pos.y))
    (time.sleep 10)
))