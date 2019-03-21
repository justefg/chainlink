import React, { Component } from 'react'
import { connect } from 'react-redux'
import { bindActionCreators, Dispatch } from 'redux'
import JobRunsList from '../../components/JobRunsList'
import { getJobRuns } from '../../actions/jobRuns'
import { IState as IReduxState } from '../../reducers'

type IProps = {
  jobRuns?: IJobRun[],
  getJobRuns: Function
}

type IState = { }

class Index extends Component<IProps, IState> {
  constructor(props: IProps) {
    super(props)
  }

  componentDidMount() {
    this.props.getJobRuns()
  }

  render() {
    return <JobRunsList jobRuns={this.props.jobRuns} />
  }
}

const jobRunsSelector = (state: IReduxState): IJobRun[] | undefined =>
  state.jobRuns.items

const mapStateToProps = (state: IReduxState) => ({
  search: state.search.query,
  jobRuns: jobRunsSelector(state)
})

const mapDispatchToProps = (dispatch: Dispatch<any>) =>
  bindActionCreators({ getJobRuns }, dispatch)

const ConnectedIndex = connect(
  mapStateToProps,
  mapDispatchToProps
)(Index)

export default ConnectedIndex
