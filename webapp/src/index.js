import { Application } from "stimulus"
import { definitionsFromContext } from "stimulus/webpack-helpers"

import "../css/normalize.css";
import "../css/skeleton.css";
import "../css/main.css";

var Turbolinks = require("turbolinks")
Turbolinks.start()

const application = Application.start()
const context = require.context("./controllers", true, /\.js$/)
application.load(definitionsFromContext(context))


