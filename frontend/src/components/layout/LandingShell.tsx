import React from 'react'
import classes from './css/Shell.module.css'
import { Outlet } from 'react-router-dom'
import { AppShell, Burger, Group } from '@mantine/core'
import { useDisclosure } from '@mantine/hooks'
import MaxWidth from '../MaxWidth'

const LandingShell = () => {
  const [openedMenu, { toggle }] = useDisclosure();

  return (
    <AppShell
      className={classes.appShell}
      header={{ height: 60 }}
    >
      <AppShell.Header>
        <MaxWidth>
          <Group h="60">
            <div>Header</div>
            <Burger size="lg" opened={openedMenu} onClick={toggle} />
          </Group>
        </MaxWidth>
      </AppShell.Header>
      <AppShell.Main>
        <MaxWidth>
          <Outlet />
        </MaxWidth>
      </AppShell.Main>
    </AppShell>
  )
}

export default LandingShell;
