import type { GetServerSidePropsContext, NextApiRequest, NextApiResponse } from "next"
import type { NextAuthOptions as NextAuthConfig } from "next-auth"
import { getServerSession } from "next-auth"

// Read more at: https://next-auth.js.org/getting-started/typescript#module-augmentation
declare module "next-auth/jwt" {
  interface JWT {
    /** The user's role. */
    userRole?: "admin"
  }
}

const endpoint = "http://localhost:8080"

export const config = {
  providers: [
    {
      id: "ninshow",
      name: "NinShow",
      type: "oauth",
      version: "2.0",
      issuer: endpoint,
      token: endpoint + "/op/token",
      userinfo: endpoint + "/op/userinfo",
      authorization: {
        url: endpoint + "/op/authorize",
        params: {
          redirect_uri: "http://localhost:3000",
          response_type: "code",
          scope: "openid profile email",
        }
      },
      jwks_endpoint: endpoint + "/op/certs",
      checks: ["nonce", "state"],
      async profile(profile, tokens) {
        return {
          id: profile.id,
          name: profile.name,
          email: profile.email,
        }
      },
      idToken: true,
      clientId: "ninshow",
      clientSecret: "ninshow",
      client: {
        authorization_signed_response_alg: "RS256",
        id_token_signed_response_alg: "RS256",
      },
    },
  ],
  callbacks: {
    async jwt({ token }) {
      token.userRole = "admin"
      return token
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
