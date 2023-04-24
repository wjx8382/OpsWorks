<template>
    <div>
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Namespace</th>
            <th>IP</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="pod in pods" :key="pod.name">
            <td>{{ pod.name }}</td>
            <td>{{ pod.namespace }}</td>
            <td>{{ pod.ip }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </template>

  <script>
  export default {
    data() {
      return {
        pods: []
      }
    },
    created() {
      fetch('http://localhost:8080/pods')
        .then(response => response.json())
        .then(data => {
          this.pods = data.pods
        })
        .catch(error => console.error(error))
    }
  }
  </script>