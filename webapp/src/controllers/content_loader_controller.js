import { Controller } from "stimulus"

export default class extends Controller {
    connect() {
        this.loadLogin()
    }

    loadLogin() {
        fetch(this.data.get("url"))
            .then(response => response.text())
            .then(html => {
                this.element.innerHTML = html
            })
    }
}