package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (tl *taskList) agregarALista(t *task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *taskList) eliminarDeLista(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func (tl *taskList) imprimirLista() {
	fmt.Println("For loop inside a struct method:")
	for _, tarea := range tl.tasks {
		fmt.Println(tarea.nombre)
	}
}

func (tl *taskList) imprimirListaCompletados() {
	fmt.Println("For loop inside a struct method:")
	for _, tarea := range tl.tasks {
		if tarea.completado {
			fmt.Println(tarea.nombre)
		}
	}
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompletada() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	// utilizamos un apuntador sobre el struct
	t := &task{
		nombre:      "completar mi curso de go",
		descripcion: "Completar mi curso de go de platzi esta semana",
	}
	t1 := &task{
		nombre:      "completar mi curso de python",
		descripcion: "Completar mi curso de python de platzi esta semana",
	}
	fmt.Printf("%+v\n", *t)
	// los metodos accederan a los valores del los apuntadores
	t.marcarCompletada()
	t.actualizarNombre("Finalizar mi curso de go")
	t.actualizarDescripcion("Completar mi curso cuanto antes.")
	fmt.Printf("%+v\n", *t)
	// agregar tarea
	lista := &taskList{
		tasks: []*task{t1},
	}
	lista.agregarALista(t)
	fmt.Println(len(lista.tasks))
	lista.agregarALista(t)
	lista.agregarALista(t)
	lista.agregarALista(t)
	fmt.Println(len(lista.tasks))
	fmt.Println(*lista.tasks[0])
	fmt.Println(*lista.tasks[1])
	lista.eliminarDeLista(2)
	fmt.Println(len(lista.tasks))

	fmt.Println("With For clause:")
	for i := 0; i < len(lista.tasks); i++ {
		fmt.Printf("Index: %d, Nombre: %s \n", i, lista.tasks[i].nombre)
	}

	fmt.Println("With Range:")
	for index, tarea := range lista.tasks {
		fmt.Printf("Index: %d, Nombre: %s \n", index, tarea.nombre)
	}
	lista.imprimirLista()
	lista.imprimirListaCompletados()

	mapaTareas := make(map[string]*taskList)
	mapaTareas["Misael"] = lista

	tr1 := &task{
		nombre:      "Aprender Ingles",
		descripcion: "Practicar con audios en ingles",
	}
	tr2 := &task{
		nombre:      "Hacer ejercicio",
		descripcion: "Hacer natacion todos los dias.",
	}

	lista2 := &taskList{
		tasks: []*task{tr1, tr2},
	}

	mapaTareas["Kathe"] = lista2

	fmt.Println("Tareas Misa:")
	mapaTareas["Misael"].imprimirLista()
	fmt.Println("Tareas Kathe:")
	mapaTareas["Kathe"].imprimirLista()

}
