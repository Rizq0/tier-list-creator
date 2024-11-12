import React from 'react'
import classes from './css/Shell.module.css'
import { Outlet } from 'react-router-dom'
import { AppShell, Burger, Group } from '@mantine/core'
import { useDisclosure } from '@mantine/hooks'

const Shell = () => {
  const [openedMenu, { toggle }] = useDisclosure();

  return (
    <AppShell
      className={classes.appShell}
      header={{ height: 60 }}
      navbar={{
        width: 300,
        breakpoint: 'md',
        collapsed: { mobile: !openedMenu, desktop: !openedMenu },
      }}
    >
      <AppShell.Header>
        <Group h="60" px="md">
          <Burger size="lg" opened={openedMenu} onClick={toggle} />
        </Group>
      </AppShell.Header>
      <AppShell.Navbar>
        Navbar
      </AppShell.Navbar>
      <AppShell.Main>
        <Outlet />
      </AppShell.Main>
    </AppShell>
  )
}

export default Shell;
