import React from 'react'
import { RouteComponentProps, Redirect } from '@reach/router'
import { connect, MapStateToProps } from 'react-redux'
import { State as AppState } from '../../reducers'

/* eslint-disable-next-line @typescript-eslint/no-empty-interface */
interface OwnProps {}

interface StateProps {
  authenticated: boolean
}

/* eslint-disable-next-line @typescript-eslint/no-empty-interface */
interface DispatchProps {}

interface Props
  extends RouteComponentProps,
    StateProps,
    DispatchProps,
    OwnProps {}

const Private: React.FC<Props> = ({ authenticated }) => {
  return authenticated ? <></> : <Redirect to="/admin/signin" noThrow />
}

const mapStateToProps: MapStateToProps<
  StateProps,
  OwnProps,
  AppState
> = state => {
  return {
    authenticated: state.adminAuth.allowed,
  }
}

const ConnectedPrivate = connect(mapStateToProps)(Private)

export default ConnectedPrivate
