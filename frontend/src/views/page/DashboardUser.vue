<script setup lang="ts">

import { ref, onMounted } from "vue"
import { getUsers, deleteUserById } from "@/services/userService"
import type { User } from "@/types/user"

const users = ref<User[]>([])
const loadUsers = async () => {

    const data = await getUsers()
    users.value = data

}

onMounted(() => {

    loadUsers()

})

const deleteUser = async (id?: string) => {

    if (!id) return

    const confirmDelete = confirm("Delete this user?")

    if (!confirmDelete) return

    await deleteUserById(id)

    users.value = users.value.filter(u => u.id !== id)

}

</script>


<template>

    <div>

        <h2 class="mb-4">Users</h2>

        <div class="card">

            <div class="card-body">

                <table class="table table-hover">

                    <thead>

                        <tr>
                            <th>ID</th>
                            <th>Name</th>
                            <th>Email</th>
                            <th>Action</th>
                        </tr>

                    </thead>

                    <tbody>

                        <tr v-for="(user, index) in users" :key="user.id">

                            <td>{{ index + 1 }}</td>
                            <td>{{ user.name }}</td>
                            <td>{{ user.email }}</td>

                            <td>

                                <button class="btn btn-danger btn-sm" @click="deleteUser(user.id)">
                                    Delete
                                </button>

                            </td>

                        </tr>

                    </tbody>

                </table>

            </div>

        </div>

    </div>

</template>