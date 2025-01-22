import {createRouter, createWebHistory} from 'vue-router'
import {useUserStore} from '../stores/user'

const routes = [
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/auth/Login.vue'),
        beforeEnter: (to, from, next) => {
            const userStore = useUserStore()
            if (userStore.token) {
                next(userStore.role === 'staff' ? '/staff/alumni' : '/alumni/basic')
            } else {
                next()
            }
        }
    },
    {
        path: '/register',
        name: 'Register',
        component: () => import('@/views/auth/Register.vue'),
    },
    {
        path: '/alumni',
        component: () => import('@/views/alumni/AlumniLayout.vue'),
        beforeEnter: (to, from, next) => {
            const userStore = useUserStore()
            if (!userStore.token || userStore.role !== 'alumni') {
                next('/login')
            } else {
                next()
            }
        },
        children: [
            {
                path: 'basic',
                name: 'AlumniBasic',
                component: () => import('@/views/alumni/profile/BasicInfo.vue')
            },
            {
                path: 'publications',
                name: 'AlumniPublications',
                component: () => import('@/views/alumni/profile/Publication.vue')
            },
            {
                path: 'studies',
                name: 'AlumniStudies',
                component: () => import('@/views/alumni/profile/StudyHistory.vue')
            },
            {
                path: 'participation',
                name: 'AlumniParticipation',
                component: () => import('@/views/alumni/profile/Participation.vue')
            },
            {
                path: 'content',
                name: 'AlumniSideContents',
                component: () => import('@/views/alumni/content/DefaultContent.vue')
            },
            {
                path: 'work',
                name: 'AlumniWork',
                component: () => import('@/views/alumni/profile/WorkExperience.vue')
            }
        ]
    },
    {
        path: '/staff',
        component: () => import('@/views/staff/StaffLayout.vue'),
        beforeEnter: (to, from, next) => {
            const userStore = useUserStore()
            if (!userStore.token || userStore.role !== 'staff') {
                next('/login')
            } else {
                next()
            }
        },
        children: [
            {
                path: 'alumni',
                name: 'AlumniManagement',
                component: () => import('@/views/staff/alumni/AlumniManagement.vue')
            },
            {
                path: 'staffs',
                name: 'StaffManagement',
                component: () => import('@/views/staff/staff/StaffTable.vue')
            },
            {
                path: 'publications',
                name: 'PublicationDatabase',
                component: () => import('@/views/staff/alumni/PublicationDatabase.vue')
            },
            {
                path: 'contents',
                name: 'ContentManagement',
                component: () => import('@/views/staff/content/ContentTable.vue')
            }
        ]
    },

    {
        path: '/',
        redirect: to => {
            const userStore = useUserStore()
            if (!userStore.token) return '/login'
            return userStore.role === 'staff' ? '/staff/alumni' : '/alumni/basic'
        }
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'NotFound',
        component: () => import('@/views/auth/Login.vue') // TODO: 404
    }
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes
})

router.beforeEach((to, from, next) => {
    const userStore = useUserStore()

    userStore.loadTokenAndRole()

    const publicPaths = ['/login', '/register']

    if (publicPaths.includes(to.path)) {
        next()
        return
    }

    if (!userStore.token) {
        next('/login')
        return
    }

    if (to.path.startsWith('/staff') && userStore.role !== 'staff') {
        next('/alumni/basic')
        return
    }

    if (to.path.startsWith('/alumni') && userStore.role !== 'alumni') {
        next('/staff/alumni')
        return
    }

    next()
})
export default router
