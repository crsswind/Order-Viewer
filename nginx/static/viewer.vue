<template>
    <div>
        <div class="row search-box">
            <div class="col-1">Search</div>
            <input class="col-10" type="text" name="searchText" v-model="searchText" />
        </div>
        <div class="date-filters">
            <div>Created Date</div>
            <input type="date" name="fromDate" v-model="fromDate" />
            <input type="date" name="toDate" v-model="toDate" />
        </div>
        <div class="total-amount">
            Total amount: {{totalAmount}}$
        </div>
        <div class="table-container">
            <table class="table">
                <thead>
                    <tr>
                        <th class="text-left" v-for="column in columns" :key="column.name">{{column.title}}</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="row in rows" :key="row.order_name">
                        <td class="text-left" v-for="column in columns" :key="column.name">{{row[column.name]}}</td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div class="row justify-content-center">
            <div class="col-1">
                <nav aria-label="Page navigation example">
                <ul class="pagination">
                    <li class="page-item">
                    <a class="page-link" v-on:click="pageIndex = pageIndex - 1 > 0 ? (pageIndex - 1):0" aria-label="Previous">
                        <span aria-hidden="true">&laquo;</span>
                    </a>
                    </li>
                    <li class="page-item"><a class="page-link">{{pageIndex + 1}}</a></li>
                    <li class="page-item">
                    <a class="page-link" v-on:click="pageIndex = rows.length == 5 ? pageIndex + 1 : pageIndex" aria-label="Next">
                        <span aria-hidden="true">&raquo;</span>
                    </a>
                    </li>
                </ul>
                </nav>
            </div>
        </div>
    </div>
</template>

<script>
module.exports = {
    data: function() {
        return {
            columns: [
                {
                    name: "order_name",
                    title: "Order name"
                },
                {
                    name: "customer_company",
                    title: "Customer Company"
                },
                {
                    name: "customer_name",
                    title: "Customer name"
                },
                {
                    name: "order_date",
                    title: "Order date"
                },
                {
                    name: "delivered_amount",
                    title: "Delivered Amount"
                },
                {
                    name: "total_amount",
                    title: "Total Amount"
                }
            ],
            rows: [],
            searchText: "",
            pageIndex: 0,
            fromDate: null,
            toDate: null
        }
    },

    watch: {
        searchText: function(val) {
            this.refreshData();
        },

        pageIndex: function(val) {
            this.refreshData();
        },

        fromDate: function(val) {
            this.refreshData();
        },

        toDate: function(val) {
            this.refreshData();
        }
    },

    computed: {
        totalAmount: function() {
            return this.rows.map(r => r.total_amount).reduce((a, b) => a + b, 0)
        }
    },

    methods: {
        refreshData: function() {
            query = `/api/v1/orders?pageIndex=${this.pageIndex}`;
            if (this.searchText != "") query += `&searchText=${this.searchText}`;
            if (this.fromDate) query += `&fromDate=${this.fromDate}`;
            if (this.toDate) query += `&toDate=${this.toDate}`;

            self = this

            axios.get(query)
                .then(function (response) {
                    if (response.status == 200) {
                        self.rows = response.data;
                    } else {
                        console.log(response);
                    }
                })
        }
    },

    mounted: function() {
        this.$nextTick(() => this.refreshData());
    }
}
</script>

<style>

</style>
