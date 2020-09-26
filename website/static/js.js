const { createApp } = Vue;

const App = {
    data() {
        return {
            a_add: "",
            b_add: "",
            total: ""
        }
    }
}
createApp(App).mount("#add")