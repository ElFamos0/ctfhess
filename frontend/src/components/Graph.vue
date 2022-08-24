<template>
    <vue3-chart-js v-bind="{ ...lines }" />
</template>

<script>
import Vue3ChartJs from "@j-t-mcc/vue3-chartjs";
import { getRequest } from "@/requests/getRequest";
export default {
  name: "GraphComponent",
  components: {
    Vue3ChartJs
  },
  async setup() {
    const res = await getRequest('/users/graph','json')
    const lines = {
        type:'line',
        data : {
            labels: res.data.days,
            datasets: [],
        }
    }
    for (let i = 0; i < res.data.users.length; i++) {
        let color = '#'+(Math.random()*0xFFFFFF<<0).toString(16)
        lines.data.datasets.push({
            label: res.data.users[i].name,
            borderColor: color,
            backgroundColor: color,
            borderWidth: 1,
            data: res.data.users[i].graph,
        })
    }
    return {
        lines,
    };
  },
};
</script>
