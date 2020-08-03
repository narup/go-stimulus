import { Controller } from "stimulus"

export default class extends Controller {
    static targets = [ "name", "output" ]

    greet() {
        this.outputTarget.innerHTML = 'Hello ' + this.name + '. This is cool!'
    }

    get name() {
        return this.nameTarget.value
    }
}