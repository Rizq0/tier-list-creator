import { Button, Flex } from '@mantine/core';
import { Component, ErrorInfo, ReactNode } from 'react';

interface Props {
  children?: ReactNode;
}

interface State {
  hasError: boolean;
}

class ErrorBoundary extends Component<Props, State> {
  constructor(props: Props) {
    super(props);
    this.state = {
      hasError: false,
    };
  }

  public static getDerivedStateFromError(): State {
    // Update state so the next render will show the fallback UI.
    return { hasError: true };
  }

  public componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    // eslint-disable-next-line no-console
    console.error('Uncaught error:', error, errorInfo);
  }

  public render() {
    const { children } = this.props;
    const { hasError } = this.state;
    if (hasError) {
      return (
        <Flex
          align="center"
          direction="column"
          justify="center"
          w="100vw"
          h="100vh"
        >
          <h1>Sorry.. there was an unexpected error</h1>
          <Button onClick={() => window.location.replace('/')}>
            Refresh
          </Button>
        </Flex>
      );
    }

    return children;
  }
}

export default ErrorBoundary;
