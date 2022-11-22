import Layout from './Layout.svelte'
import HomeView from "./views/Home.svelte"
import LibraryView from "./views/Library.svelte"
import AuthView from "./views/Auth.svelte"
import Login from './views/Login.svelte'

function userIsAdmin() {
  //check if user is admin and returns true or false
}

const routes = [
  {
    name: '/',
    component: HomeView,
    layout: Layout
  },
  {
    name: '/library',
    component: LibraryView,
    layout: Layout
  },
  {
    name: '/auth',
    component: AuthView
  },
  {
    name: "/login",
    component: Login
  }
  // { name: 'login', component: Login, layout: PublicLayout },
  // {
  //   name: 'admin',
  //   component: AdminLayout,
  //   onlyIf: { guard: userIsAdmin, redirect: '/login' },
  //   nestedRoutes: [
  //     { name: 'index', component: AdminIndex },
  //     {
  //       name: 'employees',
  //       component: '',
  //       nestedRoutes: [
  //         { name: 'index', component: EmployeesIndex },
  //         { name: 'show/:id', component: EmployeesShow },
  //       ],
  //     },
  //   ],
  // },
]

export { routes }
