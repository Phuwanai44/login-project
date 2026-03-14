<script setup lang="ts">
import { ref } from "vue"

const stats = ref([
  { title: "Users", value: 45 },
  { title: "Orders", value: 120 },
  { title: "Products", value: 30 },
  { title: "Revenue", value: "$12,000" }
])

const users = ref([
  { id: 1, name: "John", email: "john@email.com" },
  { id: 2, name: "Anna", email: "anna@email.com" }
])
</script>

<template>

  <div>

    <h2 class="mb-4">Dashboard</h2>

    <!-- Stats -->

    <div class="row mb-4">

      <div class="col-md-3" v-for="s in stats" :key="s.title">

        <div class="card shadow-sm">

          <div class="card-body">

            <h6 class="text-muted">
              {{ s.title }}
            </h6>

            <h3>
              {{ s.value }}
            </h3>

          </div>

        </div>

      </div>

    </div>



    <!-- Recent Users -->

    <div class="card">

      <div class="card-header">
        Recent Users
      </div>

      <div class="card-body">

        <table class="table">

          <thead>

            <tr>
              <th>Name</th>
              <th>Email</th>
            </tr>

          </thead>

          <tbody>

            <tr v-for="user in users" :key="user.id">

              <td>{{ user.name }}</td>
              <td>{{ user.email }}</td>

            </tr>

          </tbody>

        </table>

      </div>

    </div>

  </div>

</template>