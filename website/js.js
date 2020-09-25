Vue.createApp({
    data() {
        return {
            a_add: 0,
            b_add: 0,
            total: 0,
        }
    },
    computed() {
        this.total = this.a_add + this.b_add
    }
}).mount('#add')