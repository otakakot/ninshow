import type { GetServerSidePropsContext, NextApiRequest, NextApiResponse } from "next"
import type { NextAuthOptions as NextAuthConfig } from "next-auth"
import { getServerSession } from "next-auth"

// Read more at: https://next-auth.js.org/getting-started/typescript#module-augmentation
declare module "next-auth/jwt" {
  interface JWT {
    idToken?: string
    /** The user's role. */
    userRole?: "admin"
  }
  interface Sessin {}
}

const endpoint = "http://localhost:8080"

export const config = {
  providers: [
    {
      id: "ninshow",
      name: "NinShow",
      type: "oauth",
      wellKnown: endpoint + "/op/.well-known/openid-configuration",
      version: "2.0",
      checks: ["nonce", "state"],
      async profile(profile) {
        return {
          id: profile.sub,
          name: profile.name,
          email: profile.email,
        }
      },
      idToken: true,
      clientId: "26bf8924-c1d9-484d-8a72-db1df2b05ccd",
      clientSecret: "ninshow",
      client: {
        token_endpoint_auth_method: "client_secret_post",
        authorization_signed_response_alg: "RS256",
        id_token_signed_response_alg: "RS256",
      },
    },
  ],
  callbacks: {
    async jwt({ token, account, profile }) {
      token.userRole = "admin"
      token.idToken = account?.id_token
      token.name = profile?.name
      token.email = profile?.email
      return token
    },
    async session({ session, token }) {   
      return {
        ...session,
        user: {
          name: token.name,
          email: token.email,
        }
      }
    },
  },
} satisfies NextAuthConfig

// Helper function to get session without passing config every time
// https://next-auth.js.org/configuration/nextjs#getserversession
export function auth(...args: [GetServerSidePropsContext["req"], GetServerSidePropsContext["res"]] | [NextApiRequest, NextApiResponse] | []) {
  return getServerSession(...args, config)
}

// We recommend doing your own environment variable validation
declare global {
  namespace NodeJS {
    export interface ProcessEnv {}
  }
}
