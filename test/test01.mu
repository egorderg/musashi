(+ 123 223.42)

---------------------------------

; symbols have some kind of parent
; could be used to determine type
(defsym win
	"symbols to define window position"
	left right top bottom)

; make object with (obj :v 2 :e 3)
; predicate object with objp
(defstruct obj
	"struct docs"
	(v 1 :ro "docs")
	(e "docs"))

(defvar current-v nil "current-v docs")

; parameters should not be writeable
(defun my-fn (x :ro y :opt)
	"docs"
	(let (local-pos 'win/left)
		(set current-v vector2(x y))
		(if (and (intp x) (intp y))
			(print "secure")
			else
			(print "unsecure"))
		(print "%d %d" current-v.x current-v.y)
		(set current-v.x 10)))

(defun my-fn-two ()
	(print "test"))

(scene my-scene :name "test"
	(vbox :gap 10
		(button :id "test" :text "Hello World")))

(1)
("test")
(current-v)
(nil)

---------------------------------

; no types, use docs
;   - describe internal functions via json file
;   - simple to generate better docs via musashi cli
; floatp, intp, stringp, structp
; allow enum, struct, var, fn only on top level
;
; musashi generates a description and a third party tool can generate code
; based on it

enum(my-enum (left right top bottom))

; docs
struct(c-obj
	(v :t vector2 :ro)
)

; docs
var("current-v docs" current-v nil)

; docs
var(my-b and(true false))

; docs
; @var x docs for x
; @var y docs for y
fn(my-fn (x :ro) (y :opt) {
	let(local-v vector2(x y))
	set(current-v vector2(x y))
	if(and(intp(x) intp(y)) {
		print("secure")
	} {
		print("unsecure")
		; return
	})
	print("%d, %d" current-v.x current-v.y)
	set(local-v.x 10)
	; return
})

scene(:name "test" {
	let((x 0) (y 0) {
		button(:text "Hello World")
	})
})
