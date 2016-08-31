/*
The ValueReference package implements functionality to operate with referece as with referent.

The typical use is to instantiate ValueReference(r) as wrapper for &v, and operate Get / Set ReferentValue as you operate on v.
It is quite helpfull when you have to operate on some variable(or structure field) by referece as interface{}, but do not want to duplicate dereferencing and type switch/assertion code.

*/
package ValueReference
