'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _constants = require('./constants');

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var DelayedItems = function DelayedItems(WrappedComponent) {
  var _class, _temp;

  return _temp = _class = function (_React$PureComponent) {
    _inherits(_class, _React$PureComponent);

    function _class(props) {
      _classCallCheck(this, _class);

      var _this = _possibleConstructorReturn(this, (_class.__proto__ || Object.getPrototypeOf(_class)).call(this, props));

      var children = _this.props.children;

      _this.state = { children: children };
      return _this;
    }

    _createClass(_class, [{
      key: 'componentWillReceiveProps',
      value: function componentWillReceiveProps(_ref) {
        var nextApplyDelays = _ref.applyDelays;
        var _props = this.props,
            applyDelays = _props.applyDelays,
            children = _props.children;

        if (applyDelays !== nextApplyDelays) {
          this.setState({
            children: nextApplyDelays ? this.applyTransitionDelays() : children
          });
        }
      }
    }, {
      key: 'applyTransitionDelays',
      value: function applyTransitionDelays() {
        var _props2 = this.props,
            children = _props2.children,
            reverse = _props2.reverse;

        var numItems = _react.Children.count(children);
        var transitionDuration = _constants.TRANSITION_DURATION_MS / 1000;
        var start = _constants.TRANSITION_SCALE_ADJUSTMENT_Y;
        var delayedChildren = _react.Children.map(children, function (child, index) {
          var itemDelayFraction = void 0;
          if (reverse) {
            itemDelayFraction = (numItems - index) / numItems;
          } else {
            itemDelayFraction = index / numItems;
          }
          var delay = (start + itemDelayFraction * (1 - start)) * transitionDuration;
          return (0, _react.cloneElement)(child, { delay: delay });
        });
        return delayedChildren;
      }
    }, {
      key: 'render',
      value: function render() {
        var _props3 = this.props,
            applyDelays = _props3.applyDelays,
            children = _props3.children,
            reverse = _props3.reverse,
            otherProps = _objectWithoutProperties(_props3, ['applyDelays', 'children', 'reverse']);

        return _react2.default.createElement(WrappedComponent, _extends({}, this.state, otherProps));
      }
    }]);

    return _class;
  }(_react2.default.PureComponent), _class.propTypes = {
    children: _propTypes2.default.node,
    applyDelays: _propTypes2.default.bool,
    reverse: _propTypes2.default.bool
  }, _temp;
};
exports.default = DelayedItems;