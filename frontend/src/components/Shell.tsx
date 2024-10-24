import React from 'react'
import { Outlet } from 'react-router-dom'
import { AppShell } from '@mantine/core'

const Shell = () => {
  return (
    <AppShell
      header={{ height: 50 }}
    >
      <AppShell.Header>
        Header
      </AppShell.Header>
      <AppShell.Main>
        <Outlet />
      </AppShell.Main>
    </AppShell>
  )
}

export default Shell;
