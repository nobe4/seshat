xattr -dr com.apple.quarantine <path to file>

goreleaser

notes pour malo

* Gerer otf

    test with https://velvetyne.fr/download/?font=karrik

* Fonte variables = chaques graisses dans le pdf 
* Pages avec un texte précis et responsive a la page sans coupure de texte
* Pages avec un texte précis et corp de texte fixe, coupure du texte possible
* Format écran
use
    func NewCustom(init *InitType) (f *Fpdf) {
        return fpdfNew(init.OrientationStr, init.UnitStr, init.SizeStr, init.FontDirStr, init.Size)
    }

